/* global updatePlot */

//// Page Controls ////

let stdsTab = document.getElementById("stdsTab");
let stdsPage = document.getElementById("stdsPage");
let statsTab = document.getElementById("statsTab");
let statsPage = document.getElementById("statsPage");
let timesTab = document.getElementById("timesTab");
let timesPage = document.getElementById("timesPage");

stdsTab.onclick = function() {
    activateTab(stdsTab, stdsPage, true);
    activateTab(statsTab, statsPage, false);
    activateTab(timesTab, timesPage, false);
    clearError();
};

statsTab.onclick = function() {
    activateTab(stdsTab, stdsPage, false);
    activateTab(statsTab, statsPage, true);
    activateTab(timesTab, timesPage, false);
    clearError();
};

timesTab.onclick = function() {
    activateTab(stdsTab, stdsPage, false);
    activateTab(statsTab, statsPage, false);
    activateTab(timesTab, timesPage, true);
    clearError();
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

let saveStats = document.getElementById("saveStats");
saveStats.onchange = function() {
    var hideable = document.getElementsByClassName("stats-hideable");
    var n = hideable.length;
    for (var i=0; i < n; i++) {
        if (saveStats.checked) {
            hideable[i].style.display = "block";
        } else {
            hideable[i].style.display = "none";
        }
    }
};

function clearError() {
    let errorSect = document.getElementById("errorSect");
    errorSect.innerHTML = "";
    errorSect.style.display = "none";
}

function showError(errorStr) {
    let errorSect = document.getElementById("errorSect");
    errorSect.innerHTML = "ERROR: " + errorStr;
    errorSect.style.display = "block";
}

function disableBtn(buttonName, toDisable) {
    let btn = document.getElementById(buttonName);
    if (toDisable == true) {
        btn.disabled = true;
    } else {
        btn.disabled = false;
    }
}

//// Server Calls ////
//// For standards plotting ////

function standardsAPI() {
    let stdsPath = document.getElementById("stdsPath").value;
    let postStr = encodeURI("stdsPath=" + stdsPath);
    if (stdsPath == "") {
        showError("No path provided for standards");
        return;
    }
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
    disableBtn("updateStds", true);
}

function handlePlotResponse(xhttp) {
    let response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        showError(response["Error"]);
    } else {
        clearError();
        updatePlot(response["Data"]);
    }
    disableBtn("updateStds", false);
}

//// For sample stats ////

function statsAPI() {
    let srcPath = document.getElementById("srcPathStats").value;
    let toSave = document.getElementById("saveStats").checked;
    let outPath = document.getElementById("outPathStats").value;
    let outFile = document.getElementById("outFileStats").value;
    let postStr = encodeURI("srcPath=" + srcPath + "&toSave=" + toSave + "&outPath=" + outPath + "&outName=" + outFile);
    if (srcPath == "") {
        showError("Source and/or output path missing");
        return;
    }
    if ((toSave == true) && (outPath == "")) {
        showError("Output path missing");
        return;
    }
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            handleStatsResponse(this);
        }
    };
    xhttp.open("POST", "/sample_stats", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send(postStr);
    disableBtn("getStats", true);
}

function handleStatsResponse(xhttp) {
    let response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        showError(response["Error"]);
    } else {
        clearError();
        let resultsSect = document.getElementById("statsResults");
        resultsSect.innerHTML = "";
        let stats = response["Stats"];
        let headers = response["Headers"];
        let rowN = stats.length;
        let colN = headers.length;
        
        let table = document.createElement("table");
        let tHeader = table.createTHead();
        let tHeaderRow = tHeader.insertRow(0);
        let tHeaderCell;
        for (var j=0; j < colN; j++) {
            tHeaderCell = tHeaderRow.insertCell(j);
            tHeaderCell.innerHTML = headers[j];
        }
        let tBody = document.createElement("tbody");
        table.appendChild(tBody);
        let tRow, tCell;
        for (var i=0; i < rowN; i++) {
            tRow = tBody.insertRow(-1);
            for (var j=0; j < colN; j++) {
                tCell = tRow.insertCell(j);
                tCell.innerHTML = stats[i][headers[j]];
            }
        }
        resultsSect.appendChild(table);
    }
    disableBtn("getStats", false);
}

//// For time usage reporting ////

function timesAPI() {
    let srcPath = document.getElementById("srcPathTimes").value;
    let postStr = encodeURI("srcPath=" + srcPath);
    if (srcPath == "") {
        showError("Source path missing");
        return;
    }
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            handleTimeUsageResponse(this);
        }
    };
    xhttp.open("POST", "/time_usage", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send(postStr);
    disableBtn("getTimes", true);
}

function handleTimeUsageResponse(xhttp) {
    let response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        showError(response["Error"]);
    } else {
        clearError();
        let resultsSect = document.getElementById("timesResults");
        resultsSect.innerHTML = "";
        let stats = response["Stats"];
        let headers = response["Headers"];
        let rowN = stats.length;
        let colN = headers.length;
        
        let table = document.createElement("table");
        let tHeader = table.createTHead();
        let tHeaderRow = tHeader.insertRow(0);
        let tHeaderCell;
        for (var j=0; j < colN; j++) {
            tHeaderCell = tHeaderRow.insertCell(j);
            tHeaderCell.innerHTML = headers[j];
        }
        let tBody = document.createElement("tbody");
        table.appendChild(tBody);
        let tRow, tCell;
        for (var i=0; i < rowN; i++) {
            tRow = tBody.insertRow(-1);
            for (var j=0; j < colN; j++) {
                tCell = tRow.insertCell(j);
                tCell.innerHTML = stats[i][headers[j]];
            }
        }
        resultsSect.appendChild(table);
    }
    disableBtn("getTimes", false);
}