import {connectRouter} from 'connected-react-router';
import {combineReducers} from 'redux';
import {users} from './user.reducer';

export const createRootReducer = (history) => combineReducers({
  router: connectRouter(history),
  user: users,
});
