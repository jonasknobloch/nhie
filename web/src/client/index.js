export class NHIEClient {
  constructor() {
    this.features = new Map();
    this.endpoint = document.querySelector('script[data-nhie]').dataset.nhie;
  }

  registerFeature(key, callback) {
    this.features.set(key, callback);
  }

  encodeFeatures(url, filter) {
    this.features.forEach((callback, key) => {
      if (filter !== undefined && !filter(key)) {
        return;
      }

      const value = callback();

      if (Array.isArray(value)) {
        value.forEach((element) => {url.searchParams.append(key, element)})
      } else {
        url.searchParams.set(key, value);
      }
    });

    return url;
  }

  limitFeatures() {
    let keys = Array.from(arguments);

    return (key) => {
      return keys.includes(key);
    }
  }

  async fetchStatement() {
    return fetch(this.encodeFeatures(new URL(this.endpoint), this.limitFeatures('statement_id', 'language', 'category')));
  }
}
