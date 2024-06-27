function displayArrayData(someArray) {
    document.getElementById('tBody').innerHTML = "";

    for (let item of someArray){
        html = createHTML(item);
        document.getElementById('tBody').insertAdjacentHTML('beforeend', html);
    }
}

function sortByAny(someArray, field) {
    // console.log("Field =", field);

    if (field != oldField) {
        someArray.sort(byFieldDown(field));
        oldField = field;
    } else {
        someArray.sort(byFieldUp(field));
        oldField = '';
    }

    displayArrayData(someArray);
}

function byFieldUp(fieldName){
    return (a, b) => a[fieldName] < b[fieldName] ? 1 : -1;
}

function byFieldDown(fieldName){
    return (a, b) => a[fieldName] > b[fieldName] ? 1 : -1;
}