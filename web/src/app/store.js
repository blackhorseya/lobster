import {createBrowserHistory} from 'history';
import {applyMiddleware, compose, createStore} from 'redux';
import {createRootReducer} from '../_reducers';
import {routerMiddleware} from 'connected-react-router';
import thunk from 'redux-thunk';

export const history = createBrowserHistory();

export default function configureStore(preloadedState) {
  const enhancer = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

  return createStore(
      createRootReducer(history),
      preloadedState,
      enhancer(
          applyMiddleware(
              routerMiddleware(history),
              thunk,
          ),
      ),
  );
}
