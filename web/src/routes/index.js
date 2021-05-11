import {Route, Switch} from 'react-router';
import React from 'react';
import {AppBar, Button, Toolbar, Typography} from '@material-ui/core';
import {connect} from 'react-redux';
import {push} from 'connected-react-router';
import {routeConstants} from '../_constants';

class Routes extends React.Component {
  render() {
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
            </Toolbar>
          </AppBar>

          <Switch>
            <Route exact path={routeConstants.Root}
                   render={() => (<div>Root Page</div>)}/>
            <Route render={() => (<div>No Match</div>)}/>
          </Switch>
        </React.Fragment>
    );
  }
}

function mapStateToProps(state) {
  return {};
}

const actionCreators = {
  push: push,
};

const connectedRoutes = connect(mapStateToProps, actionCreators)(Routes);
export {connectedRoutes as Routes};