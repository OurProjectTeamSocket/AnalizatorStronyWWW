google.charts.load('current', {'packages':['timeline']});
google.charts.setOnLoadCallback(drawChart);
function drawChart() {
  var container = document.getElementById('graph5');
  var chart = new google.visualization.Timeline(container);
  var dataTable = new google.visualization.DataTable();

  dataTable.addColumn({ type: 'string', id: 'Status' });
  dataTable.addColumn({ type: 'date', id: 'Start' });
  dataTable.addColumn({ type: 'date', id: 'End' });
  dataTable.addRows([
    [ 'Down',      new Date(2021, 08, 25, 10, 05, 00),  new Date(2021, 08, 25, 10, 15, 00) ],
    [ 'Down',  new Date(2021, 08, 25, 10, 25, 00),  new Date(2021, 08, 25, 11, 20, 00) ]]);

  chart.draw(dataTable);
}