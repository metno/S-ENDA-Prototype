import React, {Component} from "react";
import MenuItem from '@material-ui/core/MenuItem';
import MenuIcon from '@material-ui/icons/Menu';
import IconButton from "@material-ui/core/IconButton/IconButton";
import Popper from "@material-ui/core/Popper/Popper";
import Paper from "@material-ui/core/Paper/Paper";
import ClickAwayListener from "@material-ui/core/ClickAwayListener/ClickAwayListener";
import MenuList from "@material-ui/core/MenuList/MenuList";
import Grow from "@material-ui/core/Grow/Grow";

class DropDownMenu extends Component {

    state = {
        anchorEl: null,
        open: false,
        currentElem: 0,
        dropDownItems: [
            {name: "Dashboard", id: 1},
            {name: "Register dataset", id: 2}
        ]
    };

    handleClick(event, id) {
        this.setState({ anchorEl: event.currentTarget, currentElem: id });
        switch(id) {
            case 1:
                window.open("/","_self");
                break;
            case 2:
                window.open("/dataset","_self");
                break;
            default:
                window.open("/","_self");
                break;
        }
    };

    handleClose = event => {
        if (this.anchorEl.contains(event.target)) {
            return;
        }

        this.setState({ open: false });
    };

    handleToggle = () => {
        this.setState(state => ({ open: !state.open }));
    };

    render() {
        const { open } = this.state;
        const dropDownList = this.state.dropDownItems.map(item => {
            return (
                <MenuItem key={item.id} onClick={(e) => this.handleClick(e, item.id)}>{item.name}</MenuItem>
            )
        });

        return (
            <div>
                <IconButton
                    buttonRef={node => {
                        this.anchorEl = node;
                    }}
                    aria-owns={open ? 'menu-list-grow' : null}
                    aria-haspopup="true"
                    onClick={this.handleToggle}
                    color="inherit"
                    aria-label="Open drawer"
                >
                    <MenuIcon/>
                </IconButton>
                <Popper open={open} anchorEl={this.anchorEl} transition disablePortal>
                    {({ TransitionProps, placement }) => (
                        <Grow
                            {...TransitionProps}
                            id="menu-list-grow"
                            style={{ transformOrigin: placement === 'bottom' ? 'center top' : 'center bottom' }}
                        >
                            <Paper>
                                <ClickAwayListener onClickAway={this.handleClose}>
                                    <MenuList>
                                        {dropDownList}
                                    </MenuList>
                                </ClickAwayListener>
                            </Paper>
                        </Grow>
                    )}
                </Popper>
            </div>
        );
    }
}

export default DropDownMenu;