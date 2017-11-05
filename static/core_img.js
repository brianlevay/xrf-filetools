/* global Image */
// This maintains any state variables //

var state_vals = {
    fileName: null,
    img: null,
    imgHeight: 0, imgWidth: 0, mult: 1.0
};

// Root function for handling the image loading //

function handleFile(files) {
    var canvasArea = document.getElementById('canvasArea');
    var canvas = document.getElementById('img_canvas');
    var ctxCanvas = canvas.getContext('2d');
    
    state_vals["img"] = null;
    state_vals["img"] = new Image();
    state_vals["img"].onload = function() {
        var height = state_vals["img"].height;
        var width = state_vals["img"].width;
        canvasArea.style.height = height + "px";
        canvasArea.style.width = width + "px";
        canvas.height = height;
        canvas.width = width;
        state_vals["imgHeight"] = height;
        state_vals["imgWidth"] = width;
        ctxCanvas.drawImage(state_vals["img"], 0, 0, width, height);
        
        var plotBtn = document.getElementById('plotPoints');
        var largerBtn = document.getElementById('largerImg');
        var smallerBtn = document.getElementById('smallerImg');
        plotBtn.disabled = false;
        largerBtn.disabled = false;
        smallerBtn.disabled = false;
    };
    state_vals["img"].src = window.URL.createObjectURL(files[0]);
    return;
}

// This handles redrawing the image on the canvas due to changes in size or rotation //

function updateImage(sizeChange) {
    var canvasArea = document.getElementById('canvasArea');
    var canvas = document.getElementById('img_canvas');
    var ctxCanvas = canvas.getContext('2d');
    if (sizeChange == "larger") {
        state_vals["mult"] = state_vals["mult"] * (1.0/0.9);
    } else if (sizeChange == "smaller") {
        state_vals["mult"] = state_vals["mult"] * (0.9/1.0);
    }
    var newWidth = Math.round(state_vals["imgWidth"] * state_vals["mult"], 0);
    var newHeight = Math.round(newWidth * (state_vals["imgHeight"] / state_vals["imgWidth"]), 0);
    
    canvasArea.style.height = newHeight + "px";
    canvasArea.style.width = newWidth + "px";
    canvas.height = newHeight;
    canvas.width = newWidth;
    
    ctxCanvas.save();
    ctxCanvas.drawImage(state_vals["img"], 0, 0, newWidth, newHeight);
    ctxCanvas.restore();
    return;
}
