import * as React from 'react';
import * as GraphLib from 'react-d3-graph';
import * as $ from 'jquery';

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
    y: number

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

class Config {
    height: number
    width: number
    nodeHighlightBehavior: boolean
    node: Object
    link: Object
    staticGraph: boolean
}

export default class Graph extends React.Component {
    static initConfig(): Config {
        return {
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
            },
            staticGraph: false
        };
    }
   
    state = {
        nodes: new Array<Node>(),
        links: new Array<Edge>(),
        root: "",
        config: Graph.initConfig(),
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
        this.state.config = Graph.initConfig();
    }

    grammarMapper(elemRaw: any): GrammarElem {
        return new GrammarElem(elemRaw.Name, elemRaw.Level, elemRaw.Children);
    }

    grammarCallback = (data: any) => {
        const root = data.root;
        const grammar = data.grammar;
        var nodes = new Array<Node>();
        var edges = new Array<Edge>();
        const grammarElems = grammar.map(this.grammarMapper);
        grammarElems.forEach((elem : any) => { nodes.push(new Node(elem.name)); });
        grammarElems.forEach((elem : GrammarElem) => {
            if (elem.children === null) {
                return;
            }
            elem.children.forEach((name: string) => {
                edges.push(new Edge(elem.name, name));
            });
        });
        /*
        var yPos = 10;
        nodes.map((n) => {
            n.y = yPos;
            yPos += 30;
        });
        */

        this.state.nodes = nodes;
        this.state.links = edges;
        this.state.root = root;
        this.setState(this.state);
    }

    componentDidMount() {
        this.updateWindowDimensions();
    }

    updateWindowDimensions = () => {
        const width:  number     = window.innerWidth;
        const height: number     = window.innerHeight;
        this.state.config.width  = width;
        this.state.config.height = height;
        this.setState(this.state);
        window.addEventListener('resize', this.updateWindowDimensions);
    }
    
    componentWillUnmount() {
          window.removeEventListener('resize', this.updateWindowDimensions);
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
                config={this.state.config}
            />
        );
    }
}
