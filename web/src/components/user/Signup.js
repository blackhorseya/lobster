import React from 'react';
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
import {connect} from 'react-redux';

class Signup extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      id: '',
      password: '',
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSignup = this.handleSignup.bind(this);
  }

  handleChange(e) {
    const {name, value} = e.target;
    this.setState({[name]: value});
  }

  handleSignup() {
    const {id, password} = this.state;
    if (id && password) {
      this.props.signup(id, password);
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
                    <Typography variant="h6">Signup</Typography>
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
                              onClick={this.handleSignup}>Submit</Button>
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
  signup: userActions.signup,
};

const connectedSignup = connect(mapStateToProps, actionCreators)(Signup);
export {connectedSignup as Signup};