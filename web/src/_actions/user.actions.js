import {userService} from '../_services';
import {userConstants} from '../_constants';

export const userActions = {
  login,
};

function login(id, password) {
  return dispatch => {
    dispatch(request());

    userService.login(id, password).then(
        resp => {
          dispatch(success(resp.data));
          localStorage.removeItem('token');
          localStorage.setItem('token', resp.data.accessToken);
        },
        error => {
          dispatch(failure(error.toString()));
        },
    );
  };

  function request() {
    return {type: userConstants.LOGIN_REQUEST};
  }

  function success(profile) {
    return {type: userConstants.LOGIN_SUCCESS, profile};
  }

  function failure(error) {
    return {type: userConstants.LOGIN_FAILURE, error};
  }
}