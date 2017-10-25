/* global updatePlot */

//// Page Controls ////

let stdsTab = document.getElementById("stdsTab");
let stdsPage = document.getElementById("stdsPage");
let filesTab = document.getElementById("filesTab");
let filesPage = document.getElementById("filesPage");

stdsTab.onclick = function() {
    activateTab(stdsTab, stdsPage, true);
    activateTab(filesTab, filesPage, false);
};

filesTab.onclick = function() {
    activateTab(stdsTab, stdsPage, false);
    activateTab(filesTab, filesPage, true);
};

function activateTab(tabObj, pageObj, activate) {
    if (activate == true) {
        tabObj.style.backgroundColor = "black";
        tabObj.style.color = "white";
        tabObj.style.fontWeight = "bold";
        pageObj.style.display = "block";
    } else {
        tabObj.style.backgroundColor = "white";
        tabObj.style.color = "black";
        tabObj.style.fontWeight = "normal";
        pageObj.style.display = "none";
    }
}

//// Server Calls ////

function updateStds() {
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            handleResponse(this);
        }
    };
    xhttp.open("POST", "/update_stds", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    var postStr = "stdsPath=" + document.getElementById("stdsPath").value;
    xhttp.send(postStr);
}

function handleResponse(xhttp) {
    var response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + response["Error"]);
    } else {
        document.getElementById("errorSect").innerHTML = "";
        updatePlot(response["Data"]);
    }
}
