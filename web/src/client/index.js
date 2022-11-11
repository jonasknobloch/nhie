export class NHIEClient {
  constructor() {
    this.features = new Map();
  }

  registerFeature(key, callback) {
    this.features.set(key, callback);
  }

  async fetchStatement() {
    let url = new URL('http://api.nhie.local/v2/statements/next')

    this.features.forEach((callback, key) => {
      const value = callback();

      if (Array.isArray(value)) {
        value.forEach((element) => {url.searchParams.append(key, element)})
      } else {
        url.searchParams.set(key, value);
      }
    })

    return fetch(url)
  }
}
