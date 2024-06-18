var stop = false;
var portMap = {};

function delRow(id) {
    document.getElementById(id).innerHTML = "";
}

function stopScan() {
    stop = true;
}

async function scanAddr(addr) {
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

    for (let i = begin ; i <= end; i++) {

        if (stop) {
            break;
        }

        let url = '/api/port/'+addr+'/'+i;
        port = await (await fetch(url)).json();

        document.getElementById("curPort").innerHTML = "Scanning port "+i;

        if (port.State) {
            createHTML(addr, port);
        }
    }
}

function createHTML(addr, port) {

    let html = `
    <tr id="row${port.Port}">
        <td>
            <input name="name" type="text" class="form-control" value="${port.Name}">
        </td>
        <td>
            <a href="http://${addr}:${port.Port}">${port.Port}</a>
            <input name="port" type="hidden" value="${port.Port}">
            <input name="state" type="hidden" value="${port.State}">
        </td>
        <td>
            ${port.State}
        </td>
        <td>
            ${port.Watch}
        </td>
        <td>
            <a href="#" onclick="delRow('row${port.Port}')">
                <i class="bi bi-x-circle"></i>
            </a>
        </td>
    </tr>`;
    document.getElementById('tBody').insertAdjacentHTML('beforeend', html);
}

async function loadSavedPorts(addr) {
    document.getElementById('tBody').innerHTML = "";

    let url = '/api/addr/'+addr;
    portMap = await (await fetch(url)).json();

    for (let port of Object.values(portMap)){
        createHTML(addr, port);
    }
}