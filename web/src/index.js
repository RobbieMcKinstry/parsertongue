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
        this.state = {
            'nodes': [
                { id: 'Harry' }
            ],
            'links': []
        };
    }

    componentWillMount() {
        console.log("Mounting");
        const url = '/grammar';
        $.getJSON(url, (data)=> {
            const root = data.root;
            const grammar = data.grammar;
            console.log(this);
            var nodes = [];
            var edges = [];

            for(var property in grammar) {
                console.log(property);
                nodes.push({
                    'id': property,
                });
                let childrenArray = grammar[property];
                if (childrenArray === null) {
                    continue;
                }
                childrenArray.forEach((name) => {
                    edges.push({
                        'source': property,
                        'target': name
                    });
                });
            }
            console.log("Nodes");
            console.log(nodes);
            console.log("Edges");
            console.log(edges);
            this.setState({ nodes, links: edges});
        });
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
