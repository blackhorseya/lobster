import {userService} from '../_services';
import {routeConstants, userConstants} from '../_constants';
import {push} from 'connected-react-router';

export const userActions = {
  login,
  signup,
};

function login(id, password) {
  return dispatch => {
    dispatch(request());

    userService.login(id, password).then(
        resp => {
          dispatch(success(resp.data));
          localStorage.removeItem('token');
          localStorage.setItem('token', resp.data.accessToken);
          dispatch(push(routeConstants.Root));
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

function signup(id, password) {
  return dispatch => {
    dispatch(request());

    userService.signup(id, password).then(
        resp => {
          dispatch(success(resp.data));
          dispatch(push(routeConstants.Root));
        },
        error => {
          dispatch(failure(error.toString()));
        },
    );
  };

  function request() {
    return {type: userConstants.SIGNUP_REQUEST};
  }

  function success(profile) {
    return {type: userConstants.SIGNUP_SUCCESS, profile};
  }

  function failure(error) {
    return {type: userConstants.SIGNUP_FAILURE, error};
  }
}