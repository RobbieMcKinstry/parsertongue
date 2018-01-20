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
                {id: 'Harry'},
                {id: 'Sally'},
                {id: 'Alice'}
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
            for(var property in grammar) {
                nodes.push({
                    'id': property
                });
            }
            console.log("Nodes " + nodes);
            this.setState({ nodes, links: []});
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
