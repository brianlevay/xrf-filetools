/* global d3 */

let standards = {
    data: [],
    plot_kVp: 9,
    plot_filter: "No-Filter",
    plot_mA: 0.25,
    plot_tmin: undefined,
    plot_tmax: undefined,
    plot_source: "Al_Ka",
    plot_data: "Area",
    vals_kVp: {},
    vals_mA: {}
};

function processData() {
    let vals_tmin = new Date(3000, 0, 1);
    let vals_tmax = new Date(1000, 0, 1);
    standards.data.forEach(function(d){
        d["SPE"]["Date"] = new Date(d["SPE"]["Date"]);
        if (d["SPE"]["Date"] < vals_tmin) {
            vals_tmin = d["SPE"]["Date"];
        }
        if (d["SPE"]["Date"] > vals_tmax) {
            vals_tmax = d["SPE"]["Date"];
        }
        standards.vals_kVp[d["SPE"]["Voltage"]] = true;
        standards.vals_mA[d["SPE"]["Current"]] = true;
    });
    if (standards.plot_tmin === undefined) {
        standards.plot_tmin = new Date(vals_tmin.valueOf());
    }
    if (standards.plot_tmax === undefined) {
        standards.plot_tmax = new Date(vals_tmax.valueOf());
    }
    return;
}

function filterPoint(d) {
    let xPass = false;
    let sizePass = false;
    let excitePass = false;
    let datePass = false;
    let meta = d["SPE"];
    if ((meta["X"] == 50) || (meta["X"] == 100) || (meta["X"] == 150)) {
        xPass = true;
    }
    if ((meta["DC"] == 10) && (meta["CC"] == 12)) {
        sizePass = true;
    }
    if ((meta["Voltage"] === standards.plot_kVp) && 
        (meta["Filter"] === standards.plot_filter) && 
        (meta["Current"] === standards.plot_mA)) {
        excitePass = true;
    }
    if ((meta["Date"] >= standards.plot_tmin) && (meta["Date"] <= standards.plot_tmax)) {
        datePass = true;
    }
    return xPass && sizePass && excitePass && datePass;
}

function xValue(d) {
    return d["SPE"]["Date"];
}

function yValue(d) {
    if (standards.plot_source === "CPS") {
        return d["SPE"]["CPS"];
    } else if (standards.plot_source === "Gain") {
        return d["Gain"];
    } else if (standards.plot_source === "Offset") {
        return d["Offset"];
    } else if (standards.plot_source === "R2") {
        return d["R2"];
    } else {
        if(d["Lines"].hasOwnProperty(standards.plot_source)) {
            return d["Lines"][standards.plot_source][standards.plot_data];
        } else {
            return 0;
        }
    }
}

function yLabel() {
    if (standards.plot_source === "CPS") {
        return "CPS";
    } else if (standards.plot_source === "Gain") {
        return "Gain";
    } else if (standards.plot_source === "Offset") {
        return "Offset";
    } else if (standards.plot_source === "R2") {
        return "R2";
    } else {
        return standards.plot_source + " " + standards.plot_data;
    }
}

// Functions for updating the plot filters and initializing the UI //

function updateUI() {
    let list_kVp = document.getElementById("list_kVp");
    let list_filter = document.getElementById("list_filter");
    let list_mA = document.getElementById("list_mA");
    let start_yr = document.getElementById("start_yr");
    let start_mo = document.getElementById("start_mo");
    let start_day = document.getElementById("start_day");
    let end_yr = document.getElementById("end_yr");
    let end_mo = document.getElementById("end_mo");
    let end_day = document.getElementById("end_day");
    let list_source = document.getElementById("list_source");
    let list_data = document.getElementById("list_data");
    
    removeOptions(list_kVp);
    removeOptions(list_mA);
    for (let val in standards.vals_kVp) {
        let option = document.createElement("option");
        option.text = val;
        option.value = val;
        list_kVp.add(option);
    }
    for (let val in standards.vals_mA) {
        let option = document.createElement("option");
        option.text = val;
        option.value = val;
        list_mA.add(option);
    }
    list_kVp.value = standards.plot_kVp;
    list_filter.value = standards.plot_filter;
    list_mA.value = standards.plot_mA;
    list_source.value = standards.plot_source;
    list_data.value = standards.plot_data;
    
    start_yr.value = standards.plot_tmin.getFullYear();
    start_mo.value = standards.plot_tmin.getMonth();
    start_day.value = standards.plot_tmin.getDate();
    end_yr.value = standards.plot_tmax.getFullYear();
    end_mo.value = standards.plot_tmax.getMonth();
    end_day.value = standards.plot_tmax.getDate();
    return;
}

function removeOptions(select) {
    for(var i = select.options.length - 1 ; i >= 0 ; i--) {
        select.remove(i);
    }
}

function updateFilters() {
    let list_kVp = document.getElementById("list_kVp");
    let list_filter = document.getElementById("list_filter");
    let list_mA = document.getElementById("list_mA");
    let start_yr = document.getElementById("start_yr");
    let start_mo = document.getElementById("start_mo");
    let start_day = document.getElementById("start_day");
    let end_yr = document.getElementById("end_yr");
    let end_mo = document.getElementById("end_mo");
    let end_day = document.getElementById("end_day");
    let list_source = document.getElementById("list_source");
    let list_data = document.getElementById("list_data");
    
    standards.plot_kVp = (+list_kVp.value);
    standards.plot_filter = list_filter.value;
    standards.plot_mA = (+list_mA.value);
    standards.plot_source = list_source.value;
    standards.plot_data = list_data.value;
    
    console.log(standards.plot_tmin, standards.plot_tmax);
    let start_yr_val = (+start_yr.value);
    let start_mo_val = (+start_mo.value);
    let start_day_val = (+start_day.value);
    let end_yr_val = (+end_yr.value);
    let end_mo_val = (+end_mo.value);
    let end_day_val = (+end_day.value);
    if (!isNaN(start_yr_val) && !isNaN(start_day_val)){
        standards.plot_tmin.setFullYear(start_yr_val);
        standards.plot_tmin.setMonth(start_mo_val);
        standards.plot_tmin.setDate(start_day_val);
        standards.plot_tmin.setHours(0, 0, 0);
    } else {
        alert("Not a valid number for starting year and/or day");
    }
    if (!isNaN(end_yr_val) && !isNaN(end_day_val)){
        standards.plot_tmax.setFullYear(end_yr_val);
        standards.plot_tmax.setMonth(end_mo_val);
        standards.plot_tmax.setDate(end_day_val);
        standards.plot_tmax.setHours(23,59,59);
    } else {
        alert("Not a valid number for ending year and/or day");
    }
    console.log(standards.plot_tmin, standards.plot_tmax);
    createPlot();
    return;
}

// Chart setup //

let SVGwidth = 1200,
SVGheight = 800,
margin = {top: 100, right: 150, bottom: 100, left: 100},
width = SVGwidth - margin.left - margin.right,
height = SVGheight - margin.top - margin.bottom;

let viewBox = "0 0 " + SVGwidth.toString() + " " + SVGheight.toString();

let svg = d3.select("#plotSect").append("svg")
    .attr("id", "plot")
    .attr("viewBox", viewBox)
    .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");
  
let tooltip = d3.select("#plotSect").append("div")
    .attr("class", "tooltip")
    .style("opacity", 0);
    
let xScaleInit = d3.scaleTime().range([0, width]),
    xAxisInit = d3.axisBottom(xScaleInit);
    
let yScaleInit = d3.scaleLinear().range([height, 0]),
    yAxisInit = d3.axisLeft(yScaleInit);
        
svg.append("g")
    .attr("class", "xAxis")
    .attr("transform", "translate(0," + height + ")")
    .call(xAxisInit);
    
svg.append("text")
    .attr("class", "xAxisLabel")
    .attr("transform", "translate(" + (width/2) + " ," + (height + 50) + ")")
    .style("text-anchor", "middle")
    .text("Measurement Date");
        
svg.append("g")
    .attr("class", "yAxis")
    .call(yAxisInit);
        
svg.append("text")
    .attr("class", "yAxisLabel")
    .attr("transform", "rotate(-90)")
    .attr("x", 0 - (height/2))
    .attr("y", 0 - Math.round((margin.left/1.5)))
    .style("text-anchor", "middle")
    .text("Throughput (CPS)");
        
let series = [50,100,150];
let legendY = height - 100;
let legend = svg.selectAll(".legend")
    .data(series)
    .enter().append("g")
    .attr("class", "legend")
    .attr("transform", function(d,i) { return "translate(0," + (legendY + (i * 25)) + ")"; });

legend.append("rect")
    .attr("x", width-20)
    .attr("width", 20)
    .attr("height", 20)
    .style("fill", color);
    
legend.append("text")
    .attr("x", width-25)
    .attr("y", 10)
    .attr("dy", "0.5em")
    .style("text-anchor", "end")
    .text(function(d) { return d; });

function color(x) {
    if (x == 50) {
        return "steelblue";
    } else if (x == 100) {
        return "orange";
    } else if (x == 150) {
        return "green";
    }
    return "black"; 
}

function tooltipHTML(d) {
    let x = xValue(d);
    let y = yValue(d);
    let dateStr = (x.getMonth() + 1) + "/" + (x.getDate()) + "/" + (x.getFullYear());
    let html = d["SPE"]["Sample"] + "<br/>X = " + d["SPE"]["X"] + "<br/>";
    html = html + dateStr + "<br/>" + y;
    return html;
}

function createPlot() {
    let filteredData = standards.data.filter(filterPoint);
    
    let xScale = d3.scaleTime().range([0, width]);
    let xAxis = d3.axisBottom(xScale);
    let xMap = function(d) { return xScale(xValue(d)); };
    xScale.domain([standards.plot_tmin, standards.plot_tmax]);
    
    let yMin = 0;
    let yMax = 0;
    let yVal = 0;
    filteredData.forEach(function(d){
        yVal = yValue(d);
        if (yVal < yMin) {
            yMin = yVal;
        }
        if (yVal > yMax) {
            yMax = yVal;
        }
    });
    let yScale = d3.scaleLinear().range([height, 0]);
    let yAxis = d3.axisLeft(yScale);
    let yMap = function(d) { return yScale(yValue(d)); };
    yScale.domain([0.9*yMin, 1.1*yMax]);
    
    let plotSect = document.getElementById("plotSect");
    
    var pts = svg.selectAll(".point").remove();
    pts = svg.selectAll(".point").data(filteredData);
    
    pts.enter().append("circle")
        .attr("class", "point")
        .attr("r", 5)
        .attr("cx", xMap)
        .attr("cy", yMap)
        .style("fill", function(d) { return color(d["SPE"]["X"]); })
        .on("mouseover", function(d) {
            tooltip.html(tooltipHTML(d))
                .style("left", (d3.event.pageX + 5) + "px")
                .style("top", (d3.event.pageY - Math.round(plotSect.offsetTop)) + "px")
                .style("opacity", 0.9);
        })
        .on("mouseout", function(d) {
            tooltip.style("opacity", 0);
        });
        
    svg.select(".xAxis")
        .attr("transform", "translate(0," + height + ")")
        .call(xAxis);
    
    svg.select(".yAxis")
        .call(yAxis);
        
    svg.select(".yAxisLabel")
        .text(yLabel);
}
    
