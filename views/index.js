dataWindObj = document.getElementById('data-wind').innerHTML
dataWind = parseInt(dataWindObj)
windStatus = ""

if (dataWind <= 6) {
    windStatus = "Aman"
    document.getElementById('wind-status').innerHTML = windStatus    
}

else if (dataWind >= 7 && dataWind <= 15) {
    windStatus = "Siaga"
    document.getElementById('wind-status').innerHTML = windStatus
    document.getElementById('wind').className = "background-animation-siaga"
}
else if (dataWind > 15 ) {
    windStatus = "Bahaya"
    document.getElementById('wind-status').innerHTML = windStatus
    document.getElementById('wind').className = "background-animation-bahaya"
}


// Water
dataWaterObj = document.getElementById('data-water').innerHTML
dataWater = parseInt(dataWaterObj)
waterStatus = ""

if (dataWater <= 5) {
    waterStatus = "Aman"
    document.getElementById('water-status').innerHTML = waterStatus    
}

else if (dataWater >= 6 && dataWater <= 8) {
    waterStatus = "Siaga"
    document.getElementById('water-status').innerHTML = waterStatus
    document.getElementById('water').className = "background-animation-siaga"
}
else if (dataWater > 8 ) {
    waterStatus = "Bahaya"
    document.getElementById('water-status').innerHTML = waterStatus
    document.getElementById('water').className = "background-animation-bahaya"
}

// Wind and Water
if ((windStatus == "Aman") && (waterStatus == "Aman")) {
    document.getElementById('all-status').innerHTML = "Aman"
    document.getElementById('circle').style.outline = "#4CAF50 solid 10px"
}
else if ((windStatus == "Aman") && (waterStatus == "Siaga")) {
    document.getElementById('all-status').innerHTML = "Siaga Banjir"   
    document.getElementById('circle').className = "circle-animation-siaga"
}
else if ((windStatus == "Aman" || windStatus == "Siaga") && (waterStatus == "Bahaya")) {
    console.log("Bahaya Banjir")   
    document.getElementById('all-status').innerHTML = "Bahaya Banjir" 
    document.getElementById('circle').className = "circle-animation-bahaya"   
}
else if ((windStatus == "Siaga") && (waterStatus == "Aman")) {
    document.getElementById('all-status').innerHTML = "Siaga Badai"  
    document.getElementById('circle').className = "circle-animation-siaga"  
}

else if ((waterStatus == "Aman" || waterStatus == "Siaga") && (windStatus == "Bahaya")) {
    document.getElementById('all-status').innerHTML = "Bahaya Badai"
    document.getElementById('circle').className = "circle-animation-bahaya"
}
else if ((windStatus == "Bahaya") && (waterStatus == "Bahaya")) {
    document.getElementById('all-status').innerHTML = "Bahaya Topan"
    document.getElementById('circle').className = "circle-animation-bahaya"
}

else if ((windStatus == "Siaga") && (waterStatus == "Siaga")) {
    document.getElementById('all-status').innerHTML = "Siaga Topan"
    document.getElementById('circle').className = "circle-animation-siaga"
}