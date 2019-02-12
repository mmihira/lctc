import * as types from 'actions/types';
import reducerCreator, { DEFAULT } from 'lib/redux/reducerCreator';

const initialData = {
  tweets: []
};

const max = 20;

const handlers = {
  [types.SOCKET_DATA]: (state, action) => {
    const d = action.data;
    return {
      tweets: d.slice(0, max).filter((v, i, a) => a.indexOf(v) === i)
    };
  },
  [DEFAULT]: state => state
};

export default reducerCreator(initialData, handlers);
