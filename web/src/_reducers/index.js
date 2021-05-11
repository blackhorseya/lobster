import {connectRouter} from 'connected-react-router';
import {combineReducers} from 'redux';

export const createRootReducer = (history) => combineReducers({
  router: connectRouter(history),
});
