google.charts.load('current', {'packages':['corechart']});
      google.charts.setOnLoadCallback(drawChart3);

      function drawChart3() {
        var data = google.visualization.arrayToDataTable([
            ['DateTime', 'Response Time'],
            ['1AM',  1000],
            ['2AM',  1170],
            ['3AM',  660],
            ['4AM',  1030]
          ]);
  
          var options = {
            hAxis: {title: 'DateTime',  titleTextStyle: {color: '#333'}},
            vAxis: {minValue: 0}
          };
  
          var chart = new google.visualization.AreaChart(document.getElementById('graph3'));
          chart.draw(data, options);
      }