import './style.css';

export class NHIELinearProgress {
    constructor() {
        this.rootElement = document.querySelector('.nhie-linear-progress');
    }

    on() {
        this.rootElement.classList.remove( 'nhie-linear-progress--disabled');
    }

    off() {
        this.rootElement.classList.add('nhie-linear-progress--disabled');
    }
}