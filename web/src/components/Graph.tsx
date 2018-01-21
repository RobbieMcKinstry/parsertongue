import * as React from 'react';
import * as GraphLib from 'react-d3-graph';
import * as $ from 'jquery';

let config = {
    height: 600,
    width: 1424,
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

interface Props {}
class Node {
    id: string

    constructor(id: string) {
        this.id = id;
    }
}
class Edge {
    source: string
    target: string

    constructor(source: string, target: string) {
        this.source = source;
        this.target = target;
    }
}

export default class Graph extends React.Component {
    
    state = {
        nodes: new Array<Node>(),
        links: new Array<Edge>(),
    }
    constructor(props: Props) {
        super(props);
        this.state.nodes = [
            new Node("Harry"),
            new Node("Sally"),
            new Node("Alice"),
        ];
        this.state.links = [
            new Edge("Harry", "Sally"),
            new Edge("Sally", "Alice"),
            new Edge("Alice", "Harry"),
        ];
    }

    grammarCallback = (data: any) => {
        const root = data.root;
        const grammar = data.grammar;
        var nodes = new Array<Node>();
        var edges = new Array<Edge>();
        Object.keys(grammar).forEach((key: string) => {
            nodes.push(new Node(key));
            var children: Array<string> = grammar[key];
            if (children === null) {
                return;
            }
            children.forEach((name: string) => {
                edges.push(new Edge(key, name));
            });
        });
        this.state.nodes = nodes;
        this.state.links = edges;
        this.setState(this.state);
    }

    componentWillMount() {
        const url = '/grammar';
        $.getJSON(url, this.grammarCallback);
    }

    render() {
        return (
            <GraphLib.Graph
                id='graph-id'
                data={this.state}
                config={config}
            />
        );
    }
}
