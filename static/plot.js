/* global d3 */

function filterPoints(d) {
    let xPass = false;
    let excitePass = false;
    if ((d["X"] == 50) || (d["X"] == 100) || (d["X"] == 150)) {
        xPass = true;
    }
    if ((d["kVp"] == 9) && (d["mA"] == 0.25) && (d["DC Slit"] == 10) && (d["CC Slit"] == 12)) {
        excitePass = true;
    }
    return xPass && excitePass;
}

// 08-01-2017 16:15:15 //
function parseDate(dateStr) {
    let dateParts = dateStr.split(" ");
    let calDate = dateParts[0];
    let timeDate = dateParts[1];
    let calParts = calDate.split("-");
    let timeParts = timeDate.split(":");
    let year = +calParts[2];
    let month = +calParts[0] - 1;
    let day = +calParts[1];
    let hour = +timeParts[0];
    let minute = +timeParts[1];
    let second = +timeParts[2];
    let dateObj = new Date(year, month, day, hour, minute, second);
    return dateObj;
}

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
    let dateStr = (d["Date"].getMonth() + 1) + "/" + (d["Date"].getDate()) + "/" + (d["Date"].getFullYear());
    let html = d["Name"] + "<br/>X = " + d["X"] + "<br/>";
    html = html + dateStr + "<br/>" + d["CPS"];
    return html;
}

//// D3 code modified from http://bl.ocks.org/weiglemc/6185069 ////
//// Modified for D3 v4 ////

d3.csv("summaries.csv", function(error, data) {
    if (error) {
        let errorDiv = document.getElementById("errorSect");
        errorDiv.innerHTML = "There was an error reading the data file.";
        return;
    }
    
    let SVGwidth = 1200,
    SVGheight = 800,
    margin = {top: 100, right: 150, bottom: 100, left: 100},
    width = SVGwidth - margin.left - margin.right,
    height = SVGheight - margin.top - margin.bottom;
    
    let svg = d3.select("#plotSect").append("svg")
        .attr("width", SVGwidth)
        .attr("height", SVGheight)
        .append("g")
        .attr("transform", "translate(" + margin.left + "," + margin.top + ")");
        
    let tooltip = d3.select("#plotSect").append("div")
        .attr("class", "tooltip")
        .style("opacity", 0);
    
    let xValue = function(d) { return d["Date"]; },
        xScale = d3.scaleTime().range([0, width]),
        xMap = function(d) { return xScale(xValue(d)); },
        xAxis = d3.axisBottom(xScale);
    
    let yValue = function(d) { return d["CPS"]; },
        yScale = d3.scaleLinear().range([height, 0]),
        yMap = function(d) { return yScale(yValue(d)); },
        yAxis = d3.axisLeft(yScale);
        
   // Name,X,Date,CPS,kVp,mA,DC Slit,CC Slit
    data.forEach(function(d){
        d["X"] = +d["X"];
        d["Date"] = parseDate(d["Date"]);
        d["CPS"] = +d["CPS"];
        d["kVp"] = +d["kVp"];
        d["mA"] = +d["mA"];
        d["DC Slit"] = +d["DC Slit"];
        d["CC Slit"] = +d["CC Slit"];
    });
    let filteredData = data.filter(filterPoints);
    xScale.domain([d3.min(data, xValue), d3.max(data, xValue)]);
    yScale.domain([0, (d3.max(data, yValue)*1.1)]);
    
    svg.append("g")
        .attr("class", "xAxis")
        .attr("transform", "translate(0," + height + ")")
        .call(xAxis);
    
    svg.append("text")
        .attr("class", "axisLabel")
        .attr("transform", "translate(" + (width/2) + " ," + (height + 50) + ")")
        .style("text-anchor", "middle")
        .text("Date");
        
    svg.append("g")
        .attr("class", "yAxis")
        .call(yAxis);
        
    svg.append("text")
        .attr("class", "axisLabel")
        .attr("transform", "rotate(-90)")
        .attr("x", 0 - (height/2))
        .attr("y", 0 - Math.round((margin.left/1.5)))
        .style("text-anchor", "middle")
        .text("CPS");
    
    svg.append("text")
        .attr("class", "chartTitle")
        .attr("transform", "translate(" + (width/2) + ", " + (-20) + ")")
        .style("text-anchor", "middle")
        .text("Standards Intensity Through Time");
        
    svg.selectAll(".point")
        .data(filteredData)
        .enter().append("circle")
        .attr("class", "point")
        .attr("r", 5)
        .attr("cx", xMap)
        .attr("cy", yMap)
        .style("fill", function(d) { return color(d["X"]); })
        .on("mouseover", function(d) {
            tooltip.html(tooltipHTML(d))
                .style("left", (d3.event.pageX + 20) + "px")
                .style("top", (d3.event.pageY - 30) + "px")
                .style("opacity", 0.9);
        })
        .on("mouseout", function(d) {
            tooltip.style("opacity", 0);
        });
    
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
        
    return;
});