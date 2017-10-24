function refreshStds() {
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            seeResponse(this);
        }
    };
    xhttp.open("POST", "/refresh_stds", true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    var postStr = "stdsPath=" + document.getElementById("stdsPath").value;
    xhttp.send(postStr);
}

function updateStds() {
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            seeResponse(this);
        }
    };
    xhttp.open("GET", "/update_stds", true);
    xhttp.send();
}

function currentStds() {
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            seeResponse(this);
        }
    };
    xhttp.open("GET", "/current_stds", true);
    xhttp.send();
}

function seeResponse(xhttp) {
    var response = JSON.parse(xhttp.response);
    if (response["Error"] != "none") {
        document.getElementById("errorSect").innerHTML = ("ERROR: " + response["Error"]);
    } else {
        document.getElementById("errorSect").innerHTML = ("DATA: ") + response["Data"][0];
    }
}
