/* global Image */
// This maintains any state variables //

var img_state = {
    fileName: null,
    img: null,
    imgHeight: 0, imgWidth: 0, mult: 1.0
};

// Root function for handling the image loading //

function handleFile(files) {
    var canvas = document.getElementById('img_canvas');
    var ctxCanvas = canvas.getContext('2d');
    
    img_state["img"] = null;
    img_state["img"] = new Image();
    img_state["img"].onload = function() {
        var height = img_state["img"].height;
        var width = img_state["img"].width;
        canvas.style.display = "block";
        canvas.style.height = height + "px";
        canvas.style.width = width + "px";
        canvas.height = height;
        canvas.width = width;
        img_state["imgHeight"] = height;
        img_state["imgWidth"] = width;
        ctxCanvas.drawImage(img_state["img"], 0, 0, width, height);
        
        var plotBtn = document.getElementById('plotPoints');
        var largerBtn = document.getElementById('largerImg');
        var smallerBtn = document.getElementById('smallerImg');
        plotBtn.disabled = false;
        largerBtn.disabled = false;
        smallerBtn.disabled = false;
    };
    img_state["img"].src = window.URL.createObjectURL(files[0]);
    return;
}

// This handles redrawing the image on the canvas due to changes in size or rotation //

function updateImage(sizeChange) {
    var canvas = document.getElementById('img_canvas');
    if (sizeChange == "larger") {
        img_state["mult"] = img_state["mult"] * (1.0/0.9);
    } else if (sizeChange == "smaller") {
        img_state["mult"] = img_state["mult"] * (0.9/1.0);
    }
    var newWidth = Math.round(img_state["imgWidth"] * img_state["mult"], 0);
    var newHeight = Math.round(newWidth * (img_state["imgHeight"] / img_state["imgWidth"]), 0);
    canvas.style.height = newHeight + "px";
    canvas.style.width = newWidth + "px";
    return;
}
