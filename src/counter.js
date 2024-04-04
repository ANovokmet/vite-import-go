export function setupCounter(element, add) {
  const setCounter = () => {
    const count = add();
    element.innerHTML = `count is ${count}`;
  }
  element.addEventListener('click', setCounter);
}
