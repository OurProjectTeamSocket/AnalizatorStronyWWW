google.charts.load('current', {'packages':['bar']});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {
var data = google.visualization.arrayToDataTable([
    ['Date Time', 'Response TIme'],
    ['1AM ', 200],
    ['1.5AM', 100],
    ['1.10AM', 1000],
    ['1.20AM', 200],
    ['1.25AM', 200],
    ['1.30AM', 200],
    ['1.35AM', 300],
    ['1.40AM', 200],
    ['1.45AM', 400],
    ['1.50AM', 200],
    ['1.55AM', 200],
    ['2AM', 200],
]);

var options = {
    chart: {

    }
};

var chart = new google.charts.Bar(document.getElementById('graph2'));

chart.draw(data, google.charts.Bar.convertOptions(options));
}