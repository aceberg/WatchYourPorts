var addr = '';
var stop = false;
var portMap = {};
var portArray = [];
var oldField = '';

function delRow(id) {
    document.getElementById(id).innerHTML = "";
}

function stopScan() {
    stop = true;
}

async function scanAddr() {
    let begin = document.getElementById("begin").value;
    let end = document.getElementById("end").value;

    if (begin == "") {
        begin = 1
    }
    if (end == "") {
        end = 65535
    }
    let port = {};
    stop = false;

    let savedPorts = Object.keys(portMap);
    // console.log("Saved ports:", savedPorts);
    document.getElementById('stopBtn').style.visibility = "visible";

    for (let i = begin ; i <= end; i++) {

        if (stop) {
            break;
        }

        let url = '/api/port/'+addr+'/'+i;
        port = await (await fetch(url)).json();

        document.getElementById("curPort").innerHTML = "Scanning port "+i;

        if ((port.State) && (!savedPorts.includes(port.Port.toString()))) {
            html = createHTML(port, true);
            document.getElementById('tBody').insertAdjacentHTML('afterbegin', html);
        }
    }

    document.getElementById('stopBtn').style.visibility = "hidden";
}

function createHTML(port, found) {
    let state = ``;
    let checked = ``;
    let sup = ``;

    if (found) {
        sup = `&nbsp;<sup style="background-color:var(--bs-danger);">new</sup>`;
    }

    if (port.Watch) {
        checked = `checked`;

        if (port.State) {
            state = `<i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>`;
        } else {
            state = `<i class="bi bi-dash-circle-fill" style="color:var(--bs-danger);"></i>`;
        }
    } else {
        state = `<i class="bi bi-circle-fill"></i>`;
    }

    let html = `
    <tr id="row${port.Port}">
        <td>
            <input name="name" type="text" class="form-control" value="${port.Name}">
        </td>
        <td>
            <a href="http://${addr}:${port.Port}">${port.Port}</a>${sup}
            <input name="port" type="hidden" value="${port.Port}">
            <input name="state" type="hidden" value="${port.State}">
        </td>
        <td>
            ${state}
        </td>
        <td>
            <div class="form-check form-switch">
                <input class="form-check-input" type="checkbox" value="${port.Port}" name="watch" ${checked}>
            </div>
        </td>
        <td>
            <a href="#" onclick="delRow('row${port.Port}')">
                <i class="bi bi-x-circle"></i>
            </a>
        </td>
    </tr>`;
    
    return html;
}

async function loadSavedPorts(addr1) {

    addr = addr1;
    
    let url = '/api/addr/'+addr;
    portMap = await (await fetch(url)).json();
    portArray = Object.values(portMap);

    displaySavedPorts();
}

function displaySavedPorts() {
    document.getElementById('tBody').innerHTML = "";

    for (let port of portArray){
        html = createHTML(port);
        document.getElementById('tBody').insertAdjacentHTML('beforeend', html);
    }
}

function sortBy(field) {
    console.log("Field =", field);

    if (field != oldField) {
        portArray.sort(byFieldDown(field));
        oldField = field;
    } else {
        portArray.sort(byFieldUp(field));
        oldField = '';
    }

    displaySavedPorts();
}

function byFieldUp(fieldName){
    return (a, b) => a[fieldName] < b[fieldName] ? 1 : -1;
}

function byFieldDown(fieldName){
    return (a, b) => a[fieldName] > b[fieldName] ? 1 : -1;
}