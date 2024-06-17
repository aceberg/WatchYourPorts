function delRow(id) {
    document.getElementById(id).innerHTML = "";
}

async function scanAddr(addr) {
    let begin = document.getElementById("begin").value;
    let end = document.getElementById("end").value;

    console.log("A:", addr, "B:", begin, "E:", end);
    if (begin == "") {
        begin = 1
    }
    if (end == "") {
        end = 65535
    }
    // let p = {};

    // for (let i = begin ; i <= end; i++) {

        // let url = '/api/port/'+addr+'/'+end;
        let url = '/api/port/192.168.2.2/22';
        p = await (await fetch(url)).json();

        console.log("p =", p);
    // }
}