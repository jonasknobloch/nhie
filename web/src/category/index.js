import './style.css';

export class NHIECategory {
  constructor(root) {
    this.rootElement = root;
  }

  get key() {
    return this.rootElement.value;
  }

  get active() {
    return this.rootElement.checked;
  }

  toggle() {
    this.rootElement.checked = !this.rootElement.checked;
  }
}
