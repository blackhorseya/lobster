import {userConstants} from '../_constants';

const initState = {
  loading: false,
  logged: false,
  data: null,
  error: '',
};

export function users(state = initState, action) {
  switch (action.type) {
    case userConstants.LOGIN_REQUEST:
      return {
        ...state,
        loading: true,
      };
    case userConstants.LOGIN_SUCCESS:
      return {
        ...state,
        loading: false,
        logged: true,
        data: action.profile,
        error: '',
      };
    case userConstants.LOGIN_FAILURE:
      return {
        ...state,
        loading: false,
        logged: false,
        data: null,
        error: action.error,
      };
    default:
      return state;
  }
}
