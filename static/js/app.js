document.addEventListener('DOMContentLoaded', function() {
     // Fetch weather data for the default location on initial load
     fetchWeather('');
});

document.getElementById('weatherForm').addEventListener('submit', function(event) {
     event.preventDefault();
     const zipCode = document.getElementById('zip').value;
     fetchWeather(zipCode);
});

function fetchWeather(zipCode) {
     const resultDiv = document.getElementById('weatherResult');

     // Show loading state
     resultDiv.innerHTML = `<p>Loading...</p>`;

     fetch(`/weather?zip=${zipCode}`)
          .then(response => {
               if (!response.ok) {
                    throw new Error('Failed to fetch weather data');
               }
               return response.json();
          })
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
          .catch(error => {
               console.error('Error fetching weather data:', error);
               resultDiv.innerHTML = `<p class="text-danger">Error fetching weather data. Please try again.</p>`;
          }); 
}