import React, {Component} from "react";
import {withStyles} from "@material-ui/core";
import PropTypes from "prop-types";
//import SampleComponent from "../../components/SampleComponent/SampleComponent";
import DataDashboardComponent from "../../components/DataDashboardComponent/DataDashboardComponent";

const styles = theme => ({
    body: {
        maxWidth: theme.spacing.getMaxWidth.maxWidth,
        margin: theme.spacing.getMaxWidth.margin,
    },
});

class Home extends Component {

    render() {
        const { classes } = this.props;
        return (
            <div className={classes.body}>
                <DataDashboardComponent/>
            </div>
        );
    }

}

Home.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Home);
