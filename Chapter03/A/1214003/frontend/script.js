let cardData = function() {
    return {
        countUp: function(target,startVal,endVal,decimals,duration){
            const countUp = new CountUp(target,startVal||0,endVal,decimals||0,duration||2);
            countUp.start();
        },
        sessions: [
            {
                "label": "Hadir",
                "size": 70,
                "color": "indigo-600"
            },
            {
                "label": "Absen",
                "size": 30,
                "color": "indigo-400"
            },
        ]
    }
}