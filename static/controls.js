/* global standards, processData, updateUI, createPlot */

//// Page Controls ////

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

//// For standards plotting ////

function getStandardsAPI() {
    let xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            handlePlotResponse(this);
        }
    };
    xhttp.open("GET", "/standards", true);
    xhttp.send();
    disableBtn("updateStds", true);
    disableBtn("updateFilters", true);
}

function handlePlotResponse(xhttp) {
    let response = JSON.parse(xhttp.response);
    if(response.hasOwnProperty("Error")) {
        showError(response["Error"]);
    } else {
        clearError();
        standards.data = response["Data"];
        processData();
        updateUI();
        createPlot();
    }
    disableBtn("updateStds", false);
    disableBtn("updateFilters", false);
}

//// Initial calls on page load ////

document.addEventListener('DOMContentLoaded', function(){ 
    getStandardsAPI();
}, false);