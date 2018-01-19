import React from 'react';
import ReactDOM from 'react-dom';
import { Graph } from 'react-d3-graph';

const data = {
    nodes: [
        {id: 'Harry'},
        {id: 'Sally'},
        {id: 'Alice'}
    ],
    links: [
        {source: 'Harry', target: 'Sally'},
        {source: 'Harry', target: 'Alice'},
    ]
};

const myConfig = {
    nodeHighlightBehavior: true,
    node: {
        color: 'lightgreen',
        size: 120,
        highlightStrokeColor: 'blue'
    },
    link: {
        highlightColor: 'lightblue'
    }
};


class App extends React.Component {
    render () {
        return (
            <Graph
                id='graph-id' // id is mandatory, if no id is defined rd3g will throw an error
                data={data}
                config={myConfig}
            />
        );
    }
}

ReactDOM.render(
  <App/>, document.getElementById('react-body')
);
