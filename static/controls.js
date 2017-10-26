/* global updatePlot */

//// Page Controls ////

let stdsTab = document.getElementById("stdsTab");
let stdsPage = document.getElementById("stdsPage");
let filesTab = document.getElementById("filesTab");
let filesPage = document.getElementById("filesPage");

stdsTab.onclick = function() {
    activateTab(stdsTab, stdsPage, true);
    activateTab(filesTab, filesPage, false);
    document.getElementById("errorSect").innerHTML = "";
};

filesTab.onclick = function() {
    activateTab(stdsTab, stdsPage, false);
    activateTab(filesTab, filesPage, true);
    document.getElementById("errorSect").innerHTML = "";
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
    let stdsPath = document.getElementById("stdsPath").value;
    let postStr = encodeURI("stdsPath=" + stdsPath);
    if (stdsPath == "") {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + "No path provided for standards");
    } else {
        var xhttp;
        xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                handlePlotResponse(this);
            }
        };
        xhttp.open("POST", "/update_stds", true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send(postStr);
    }
}

function handlePlotResponse(xhttp) {
    let response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + response["Error"]);
    } else {
        document.getElementById("errorSect").innerHTML = "";
        updatePlot(response["Data"]);
    }
}

function uniqueNames() {
    let srcPath = document.getElementById("srcPathNames").value;
    let outPath = document.getElementById("outPathNames").value;
    let outFile = document.getElementById("outFileNames").value;
    let postStr = encodeURI("srcPath=" + srcPath + "&outPath=" + outPath + "&outName=" + outFile);
    if ((srcPath == "") || (outPath == "")) {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + "Source and/or output path missing");
    } else {
        var xhttp;
        xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                handleNamesResponse(this);
            }
        };
        xhttp.open("POST", "/unique_names", true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send(postStr);
    }
}

function handleNamesResponse(xhttp) {
    let response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + response["Error"]);
    } else {
        document.getElementById("errorSect").innerHTML = "";
        alert("Finished!");
    }
}

function sortFiles() {
    let srcPath = document.getElementById("srcPathSort").value;
    let outPath = document.getElementById("outPathSort").value;
    let toRename = document.getElementById("renameCheckBox").checked;
    let matchedFile = document.getElementById("matchedFileSort").value;
    let postStr = encodeURI("srcPath=" + srcPath + "&outPath=" + outPath + "&toRename=" + toRename + "&matchedFile=" + matchedFile);
    if ((srcPath == "") || (outPath == "")) {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + "Source and/or output path missing");
        return;
    }
    if ((toRename == true) && (matchedFile == "")) {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + "No path provided for matched names file");
        return;
    }
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            handleSortResponse(this);
        }
    };
    xhttp.open("POST", "/unique_names", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send(postStr);
}

function handleSortResponse(xhttp) {
    let response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + response["Error"]);
    } else {
        document.getElementById("errorSect").innerHTML = "";
        alert("Finished!");
    }
}