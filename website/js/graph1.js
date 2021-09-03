// Load the Visualization API and the corechart package.
google.charts.load('current', {'packages':['corechart']});

// Set a callback to run when the Google Visualization API is loaded.
google.charts.setOnLoadCallback(drawChart);

// Callback that creates and populates a data table,
// instantiates the pie chart, passes in the data and
// draws it.
function drawChart() {

  // Create the data table.
  var data = new google.visualization.DataTable();
  data.addColumn('string', 'Status');
  data.addColumn('number', 'state');
  data.addRows([
    ['Site UP', 3],
    ['Site Down', 1],
  ]);

  // Set chart options
  var options = {'title':'',
                 'width':'100%',
                 'height':'300%'};

  // Instantiate and draw our chart, passing in some options.
  var chart = new google.visualization.PieChart(document.getElementById('graph1'));
  chart.draw(data, options);
}