const endpoint = `${process.env.REACT_APP_API_ENDPOINT || ''}`;

export const userService = {
  login,
  signup,
};

function login(id, password) {
  const opts = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: `email=${id}&password=${password}`,
  };

  return fetch(`${endpoint}/api/v1/auth/login`, opts).then(handleResponse);
}

function signup(id, password) {
  const opts = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: `email=${id}&password=${password}`,
  };

  return fetch(`${endpoint}/api/v1/auth/signup`, opts).then(handleResponse);
}

function handleResponse(resp) {
  return resp.json().then(body => {
    if (!resp.ok) {
      const error = (body && body.msg) || resp.statusText;
      return Promise.reject(error);
    }

    return {
      data: body.data,
    };
  });
}
