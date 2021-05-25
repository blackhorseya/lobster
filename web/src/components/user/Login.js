import React from 'react';
import {connect} from 'react-redux';
import {
  AppBar,
  Button,
  Container,
  Grid,
  Paper,
  TextField,
  Toolbar,
  Typography,
} from '@material-ui/core';
import {userActions} from '../../_actions';

class Login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      id: '',
      password: '',
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleLogin = this.handleLogin.bind(this);
  }

  handleChange(e) {
    const {name, value} = e.target;
    this.setState({[name]: value});
  }

  handleLogin() {
    const {id, password} = this.state;
    if (id && password) {
      this.props.login(id, password);
    }
  }

  render() {
    const {id, password} = this.state;

    return (
        <Container maxWidth={'xs'} style={{padding: 20}}>
          <Paper style={{padding: 20}}>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <AppBar position="static">
                  <Toolbar>
                    <Typography variant="h6">Login</Typography>
                  </Toolbar>
                </AppBar>
              </Grid>
              <Grid item xs={12}>
                <form>
                  <Grid container spacing={2}>
                    <Grid item xs={12}>
                      <TextField
                          name={'id'}
                          label="ID"
                          type="text"
                          variant="outlined"
                          fullWidth
                          value={id}
                          onChange={this.handleChange}
                      />
                    </Grid>
                    <Grid item xs={12}>
                      <TextField
                          name={'password'}
                          label="Password"
                          type="password"
                          autoComplete="current-password"
                          variant="outlined"
                          fullWidth
                          value={password}
                          onChange={this.handleChange}
                      />
                    </Grid>
                    <Grid item xs={12}>
                      <Button variant="contained" color="primary"
                              onClick={this.handleLogin}>Login</Button>
                    </Grid>
                  </Grid>
                </form>
              </Grid>
            </Grid>
          </Paper>
        </Container>
    );
  }
}

function mapStateToProps(state) {
  const {user} = state;
  return {user};
}

const actionCreators = {
  login: userActions.login,
};

const connectedLogin = connect(mapStateToProps, actionCreators)(Login);
export {connectedLogin as Login};
