const graphDiv = document.getElementById('graph');

fetch("https://sarthak-service-dssd.onrender.com")
    .then(async res => {
        const data = await res.json();

        const n = 60;

        const benchTrace = {
            x: data.age.filter((_, i) => i % n === 0),
            y: data.bench.filter((_, i) => i % n === 0),
            mode: 'markers',
            name: 'Bench',
            type: 'scatter',
        };

        const squatTrace = {
            x: data.age.filter((_, i) => i % n === 0),
            y: data.squat.filter((_, i) => i % n === 0),
            mode: 'markers',
            name: 'Squat',
            type: 'scatter',
        };

        const layout = {
            title: 'Powerlifting Stats By Age',
            xaxis: {
                title: 'Age',
            },
            yaxis: {
                title: 'Weight',
            },
        };

        Plotly.newPlot(graphDiv, [benchTrace, squatTrace], layout);
    })
    .catch(error => {
        console.error('Error fetching data:', error);
    });
