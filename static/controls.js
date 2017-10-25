/* global updatePlot */

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
