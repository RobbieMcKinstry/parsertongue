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

class GrammarElem {
    name: string
    level: number
    children: Array<string>

    constructor(name: string, level: number, children: Array<string>) {
        this.name = name;
        this.level = level;
        this.children = children;
    }
}
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
        root: "",
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
        grammar.forEach((elemRaw: any) => {
            const elem = new GrammarElem(elemRaw.Name, elemRaw.Level, elemRaw.Children);
            nodes.push(new Node(elem.name));
            var children: Array<string> = elem.children;
            if (children === null) {
                return;
            }
            children.forEach((name: string) => {
                edges.push(new Edge(elem.name, name));
            });
        });
        this.state.nodes = nodes;
        this.state.links = edges;
        this.state.root = root;
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
