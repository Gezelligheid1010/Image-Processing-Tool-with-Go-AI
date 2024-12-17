export function developmentConfig(filename) {
    let envConfig = {
      BASE_URL: "http://todo-app-barkend.herokuapp.com/todos/",
      ENV: "DEVELOPMENT",
    };
    return Object.assign({}, envConfig, filename);
  }
  
  export function stagingConfig(filename) {
    let envConfig = {
      BASE_URL: "https://todo-app-barkend.herokuapp.com/todos/",
      ENV: "STAGING",
    };
    return Object.assign({}, envConfig, filename);
  }
  