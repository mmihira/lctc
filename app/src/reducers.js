import { connectRouter } from 'connected-react-router';
import { combineReducers } from 'redux';
import layout from 'reducers/layout';
import tweets from 'reducers/tweets';

export default history =>
  combineReducers({
    router: connectRouter(history),
    layout,
    tweets
  });
