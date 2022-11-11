import './style.css';

export class NHIECategorySelection {
  constructor() {
    this.rootElement = document.querySelector('.nhie-category-selection');
    this.categories = [];
    this.stopPropagation();
  }

  stopPropagation() {
    this.rootElement.addEventListener('click', (e) => e.stopPropagation());
  }

  registerCategories(...args) {
    this.categories = this.categories.concat(Array.from(args));
    this.watchCategories();
  }

  watchCategories() {
    this.categories.forEach(function(category) {
      category.rootElement.addEventListener('change', () => this.handleCategoryChange());
    }.bind(this));
  }

  handleCategoryChange() {
    if (this.activeCategories.length === 0) {
      this.categories.forEach(function(category) {
        category.toggle();
      });
    }
  }

  get activeCategories() {
    return this.categories.filter(function(category) {
      return category.active;
    });
  }
}
