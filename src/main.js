import './style.css';
import { setupCounter } from './counter.js';
import load from './main.go';

document.querySelector('#app').innerHTML = `
  <div>
    <h1>Hello Vite!</h1>
    <div class="card">
      <button id="counter" type="button">click to count</button>
    </div>
    <p class="read-the-docs">
      Click on the Vite logo to learn more
    </p>
  </div>
`;

load().then(() => {
  setupCounter(document.querySelector('#counter'), () => add());

  const output = formatJSON('{"website":"github.com/ANovokmet","projects":[{"title":"vite-import-go","url":"github.com/ANovokmet/vite-import-go"}]}');
  console.log('output', output);
});
