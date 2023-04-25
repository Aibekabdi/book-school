import host from '../main';

// eslint-disable-next-line import/prefer-default-export
export function sendRequest(route, method, body, token = '') {
  const url = host + route;

  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open(method, url);

    xhr.setRequestHeader('Accept', 'application/json');
    xhr.setRequestHeader('Content-Type', 'application/json');
    if (token !== '') {
      xhr.setRequestHeader('Authorization', `Bearer ${token}`);
    }
    xhr.responseType = 'json';

    xhr.onload = () => {
      if (xhr.status === 200) {
        resolve(xhr.response);
      } else {
        reject(xhr.response);
      }
    };

    xhr.send(JSON.stringify(body));
  });
}
