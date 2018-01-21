import React from 'react';
import ReactDOM from 'react-dom';
import { Graph } from 'react-d3-graph';
import $ from 'jquery';

const data = {
    nodes: [
    ],
    links: [
        {source: 'Harry', target: 'Sally'},
        {source: 'Harry', target: 'Alice'},
    ]
};

const myConfig = {
    height: 600,
    width:  1424,
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

    constructor(props) {
        super(props);
        this.grammarCallback = this.grammarCallback.bind(this);
        this.state = {
            'nodes': [
                { id: 'Harry' }
            ],
            'links': []
        };
    }

    grammarCallback(data) {
        const root = data.root;
        const grammar = data.grammar;
        var nodes = [];
        var edges = [];

        Object.keys(grammar).forEach((key) => {
           nodes.push({
                id: key,
            });
            let childrenArray = grammar[key];
            if (childrenArray === null) {
                return;
            }
            childrenArray.forEach((name) => {
                edges.push({
                    'source': key,
                    'target': name
                });
            });
        });
        this.setState({ nodes, links: edges});
    }

    componentWillMount() {
        const url = '/grammar';
        $.getJSON(url, this.grammarCallback);
    }

    render() {
        return (
            <Graph
                id='graph-id'
                data={this.state}
                config={myConfig}
            />
        );
    }
}

ReactDOM.render(
  <App/>, document.getElementById('react-body')
);
