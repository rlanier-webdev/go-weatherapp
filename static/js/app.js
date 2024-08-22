document.getElementById('weatherForm').addEventListener('submit', function(event) {
     event.preventDefault();
     const zipCode = document.getElementById('zip').value;

     fetch(`/weather?zip=${zipCode}`)
          .then(response => response.json())
          .then(data => {
               // Display weather information in the 'weatherResult' div
               const resultDiv = document.getElementById('weatherResult');
               resultDiv.innerHTML = `
                    <h2>Weather in ${data.location.name}, ${data.location.region},${data.location.country}</h2>
                    <p>Temperature: ${data.current.temp_f}°F (${data.current.temp_c}°C)</p>
                    <p>Condition: ${data.current.condition.text}</p>
                    <p>Humidity: ${data.current.humidity}%</p>
                    <p>Wind: ${data.current.wind_mph} mph (${data.current.wind_kph} kph)</p>
                    <p>Feels Like: ${data.current.feelslike_f}°F (${data.current.feelslike_c}°C)</p>
               `;
          })
          .catch(error => console.error('Error fetching weather data:', error))
});