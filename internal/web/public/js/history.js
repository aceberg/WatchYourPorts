var oldField = '';
var histArray = {};

loadHistory();

function createHTML(hist) {
    
    let allState = "";
    let color = "";

    for (let state of hist.State){
        if (state.State) {
            color = `bi-check-circle-fill" style="color:var(--bs-success);"`;
        } else {
            color = `bi-dash-circle-fill" style="color:var(--bs-danger);"`;
        }
        allState = allState + `<i class="bi ${color} title="${state.Date}"></i>`;
    }

    let html = `
    <tr>
        <td><a href="/scan/?addr=${hist.Addr}">${hist.Name}</a></td>
        <td><a href="/scan/?addr=${hist.Addr}">${hist.Addr}</a></td>
        <td><a href="http://${hist.Addr}:${hist.Port}">${hist.Port}</a></td>
        <td><a href="http://${hist.Addr}:${hist.Port}">${hist.PortName}</a></td>
        <td>${allState}</td>
    </tr>`;
    
    return html;
}

async function loadHistory() {
    
    let url = '/api/history';
    let histMap = await (await fetch(url)).json();
    if (histMap != null) {
        histArray = Object.values(histMap);
    }

    displayHistory();
}

function displayHistory() {
    document.getElementById('tBody').innerHTML = "";

    for (let hist of histArray){
        html = createHTML(hist);
        document.getElementById('tBody').insertAdjacentHTML('beforeend', html);
    }
}

function sortBy(field) {
    console.log("Field =", field);

    if (field != oldField) {
        histArray.sort(byFieldDown(field));
        oldField = field;
    } else {
        histArray.sort(byFieldUp(field));
        oldField = '';
    }

    displayHistory();
}

function byFieldUp(fieldName){
    return (a, b) => a[fieldName] < b[fieldName] ? 1 : -1;
}

function byFieldDown(fieldName){
    return (a, b) => a[fieldName] > b[fieldName] ? 1 : -1;
}