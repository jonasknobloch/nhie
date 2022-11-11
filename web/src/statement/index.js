import './style.css';

export class NHIEStatement {
  constructor() {
    this.rootElement = document.querySelector('.nhie-statement');
  }

  set ID(string) {
    this.rootElement.setAttribute('data-id', string);
  }

  set statement(string) {
    this.rootElement.innerHTML = string;
  }

  get ID() {
    return this.rootElement.getAttribute('data-id');
  }
}
