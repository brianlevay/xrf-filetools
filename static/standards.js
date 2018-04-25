/* global d3 */

let standards = {
    data: [],
    plot_kVp: 9,
    plot_filter: "No-Filter",
    plot_mA: 0.25,
    plot_tmin: new Date(),
    plot_tmax: new Date(),
    plot_source: "CPS",
    plot_data: "Area",
    vals_kVp: {},
    vals_mA: {},
    vals_tmin: new Date(3000,0,1),
    vals_tmax: new Date(1000,0,1)
};

function processData() {
    standards.data.forEach(function(d){
        d["SPE"]["Date"] = Date.parse(d["SPE"]["Date"]);
        if (d["SPE"]["Date"] < standards.vals_tmin) {
            standards.vals_tmin = d["SPE"]["Date"];
        }
        if (d["SPE"]["Date"] > standards.vals_tmax) {
            standards.vals_tmax = d["SPE"]["Date"];
        }
        standards.vals_kVp[d["SPE"]["Voltage"]] = true;
        standards.vals_mA[d["SPE"]["Current"]] = true;
    });
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
    if ((meta["Voltage"] == standards.plot_kVp) && 
        (meta["Filter"] == standards.plot_filter) && 
        (meta["Current"] == standards.plot_mA)) {
        excitePass = true;
    }
    if ((meta["Date"] >= standards.plot_tmin) && (meta["Date"] <= standards.plot_tmax)) {
        datePass = true;
    }
    return xPass && sizePass && excitePass && datePass;
}

function getXvalue(d) {
    return d["SPE"]["Date"];
}

function getYvalue(d) {
    if (standards.plot_source == "CPS") {
        return d["SPE"]["CPS"];
    } else if (standards.plot_source == "Gain") {
        return d["Gain"];
    } else if (standards.plot_source == "Offset") {
        return d["Offset"];
    } else if (standards.plot_source == "R2") {
        return d["R2"];
    } else {
        return d["Lines"][standards.plot_source][standards.plot_data];
    }
}

// Initial chart setup //

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
    
svg.append("g")
    .attr("class", "xAxis")
    .attr("transform", "translate(0," + height + ")");
    
svg.append("text")
    .attr("class", "xAxisLabel")
    .attr("transform", "translate(" + (width/2) + " ," + (height + 50) + ")")
    .style("text-anchor", "middle")
    .text("Date");
        
svg.append("g")
    .attr("class", "yAxis");
        
svg.append("text")
    .attr("class", "yAxisLabel")
    .attr("transform", "rotate(-90)")
    .attr("x", 0 - (height/2))
    .attr("y", 0 - Math.round((margin.left/1.5)))
    .style("text-anchor", "middle")
    .text("CPS");
        
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
    let x = getXvalue(d);
    let y = getYvalue(d);
    let dateStr = (x.getMonth() + 1) + "/" + (x.getDate()) + "/" + (x.getFullYear());
    let html = d["SPE"]["Sample"] + "<br/>X = " + d["SPE"]["X"] + "<br/>";
    html = html + dateStr + "<br/>" + y;
    return html;
}

function updatePlot() {
    let filteredData = standards.data.filter(filterPoint);
    
    let xScale = d3.scaleTime().range([0, width]);
    let xAxis = d3.axisBottom(xScale);
    let xValue = function(d) { return getXvalue(d); };
    let xMap = function(d) { return xScale(xValue(d)); };
    xScale.domain([standards.plot_tmin, standards.plot_tmax]);
    
    let yScale = d3.scaleLinear().range([height, 0]);
    let yAxis = d3.axisLeft(yScale);
    let yMap = function(d) { return yScale(yValue(d)); };
    let yValue = function(d) { return getYvalue(d); };
    yScale.domain([0, (d3.max(filteredData, yValue)*1.1)]);
    
    let plotSect = document.getElementById("plotSect");
    
    var pts = svg.selectAll(".point").remove();
    pts = svg.selectAll(".point").data(filteredData);
    
    pts.enter().append("circle")
        .attr("class", "point")
        .attr("r", 5)
        .attr("cx", xMap)
        .attr("cy", yMap)
        .style("fill", function(d) { return color(d["X"]); })
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
        .text()
}
    
