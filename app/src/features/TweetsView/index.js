import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import act from 'actions/index';
import './styles/index.scss';
import TweetEmbed from 'react-tweet-embed';

class TweetsView extends React.Component {
  constructor (props) {
    super(props);
  }

  componentDidMount () {}

  render () {
    return (
      <div className="tweet-view-container">
        { this.props.tweets.tweets.map( e => (
          <div key={e}>
            <TweetEmbed id={e} width="250px" className="tweet" />
          </div>
        ))}
      </div>
    );
  }
}

const mapDispatchToProps = dispatch => {
  return {
    setTweetsViewDimensions: (width, height) =>
      dispatch(act.setTweetsViewDimensions(width, height))
  };
};

TweetsView.propTypes = {
  tweets: PropTypes.object,
  setTweetsViewDimensions: PropTypes.func
};

const mapStateToProps = state => {
  const { tweets } = state;
  return { tweets };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TweetsView);
