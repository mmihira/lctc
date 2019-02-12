import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import act from 'actions/index';
import './styles/index.scss';
import Login from 'features/Login/index';
import TweetsView from 'features/TweetsView/index';

class App extends React.Component {
  constructor (props) {
    super(props);
    this.props.setAppDimensions(window.innerWidth, window.innerHeight);
  }

  componentDidMount () {
  }

  render () {
    return (
      <div className="app-root">
        <div className="header">
          <Login />
        </div>
        <div className="body">
          <TweetsView />
        </div>
      </div>
    );
  }
}

const mapDispatchToProps = dispatch => {
  return {
    setAppDimensions: (width, height) =>
      dispatch(act.setAppDimensions(width, height)),
  };
};

App.propTypes = {
  setAppDimensions: PropTypes.func
};

export default connect(
  null,
  mapDispatchToProps
)(App);
