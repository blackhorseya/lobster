import {Route, Switch} from 'react-router';
import React from 'react';
import {
  AppBar,
  Button,
  IconButton,
  Toolbar,
  Typography,
} from '@material-ui/core';
import {connect} from 'react-redux';
import {push} from 'connected-react-router';
import {routeConstants} from '../_constants';
import {AccountCircle} from '@material-ui/icons';
import {Login} from '../components/user';

class Routes extends React.Component {
  render() {
    const {user} = this.props;

    return (
        <React.Fragment>
          <AppBar position="static">
            <Toolbar>
              <Button color={'inherit'}
                      onClick={() => this.props.push(routeConstants.Root)}>
                <Typography variant="h5">
                  Lobster
                </Typography>
              </Button>
              {user.logged === false && (
                  <Button color="inherit" onClick={() => this.props.push(
                      routeConstants.Login)}>Login</Button>
              )}
              {user.logged && (
                  <div>
                    <IconButton
                        aria-label="account of current user"
                        aria-controls="menu-appbar"
                        aria-haspopup="true"
                        color="inherit"
                    >
                      <AccountCircle/>
                    </IconButton>
                  </div>
              )}
            </Toolbar>
          </AppBar>

          <Switch>
            <Route exact path={routeConstants.Root}
                   render={() => (<div>Root Page</div>)}/>
            <Route path={routeConstants.Login} component={Login}/>
            <Route render={() => (<div>No Match</div>)}/>
          </Switch>
        </React.Fragment>
    );
  }
}

function mapStateToProps(state) {
  const {user} = state;
  return {user};
}

const actionCreators = {
  push: push,
};

const connectedRoutes = connect(mapStateToProps, actionCreators)(Routes);
export {connectedRoutes as Routes};