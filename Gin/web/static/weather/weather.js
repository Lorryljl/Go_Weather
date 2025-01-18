(document).ready(function() {
    // 点击查询按钮时触发
    ('#search-btn').click(function() {
        // 获取用户输入的城市和选择的天气类型
        var city = $('#city-input').val();
        var weatherType = $('#baseAll').val();

        if (city) {
            // 发送请求到后端获取天气数据
            $.ajax({
                url: '/getWeather',
                method: 'POST',
                contentType: 'application/json',
                data: JSON.stringify({
                    city: city,
                    weatherType: weatherType
                }),
                success: function(response) {
                    // 显示返回的天气数据
                    if (response.status == "1") {
                        // 填充基本天气信息
                        $('#city-name').text(response.forecasts[0].city);
                        $('#province-name').text(response.forecasts[0].province);
                        $('#report-time').text(response.forecasts[0].reportTime);

                        // 填充天气预报表格
                        var forecastTbody = $('#forecast-tbody');
                        forecastTbody.empty(); // 清空现有数据

                        response.forecasts[0].casts.forEach(function(cast) {
                            var row = $('<tr>');
                            row.append('<td>' + cast.date + '</td>');
                            row.append('<td>' + cast.week + '</td>');
                            row.append('<td>' + cast.dayWeather + '</td>');
                            row.append('<td>' + cast.nightWeather + '</td>');
                            row.append('<td>' + cast.dayTemp + '°C</td>');
                            row.append('<td>' + cast.nightTemp + '°C</td>');
                            row.append('<td>' + cast.dayWind + '</td>');
                            row.append('<td>' + cast.nightWind + '</td>');
                            row.append('<td>' + cast.dayPower + '</td>');
                            row.append('<td>' + cast.nightPower + '</td>');
                            forecastTbody.append(row);
                        });
                    } else {
                        alert('获取天气信息失败: ' + response.info);
                    }
                },
                error: function() {
                    alert('请求失败，请稍后再试');
                }
            });
        } else {
            alert('请输入城市名称');
        }
    });
});