package ui

import (
	"io/fs"

	"github.com/knipferrc/fm/constants"
	"github.com/knipferrc/fm/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type directoryMsg []fs.FileInfo
type fileContentMsg string

func updateDirectoryListing(dir string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		files := utils.GetDirectoryListing(dir, showHidden)

		return directoryMsg(files)
	}
}

func renameFileOrDir(filename, value string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		utils.RenameDirOrFile(filename, value)
		files := utils.GetDirectoryListing(constants.CurrentDirectory, showHidden)

		return directoryMsg(files)
	}
}

func moveDir(dir, value string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		utils.CopyDir(dir, value, true)
		files := utils.GetDirectoryListing(constants.CurrentDirectory, showHidden)

		return directoryMsg(files)
	}
}

func moveFile(file, value string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		utils.CopyFile(file, value, true)
		files := utils.GetDirectoryListing(constants.CurrentDirectory, showHidden)

		return directoryMsg(files)
	}
}

func deleteDir(dir string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		utils.DeleteDirectory(dir)
		files := utils.GetDirectoryListing(constants.CurrentDirectory, showHidden)

		return directoryMsg(files)
	}
}

func deleteFile(file string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		utils.DeleteFile(file)
		files := utils.GetDirectoryListing(constants.CurrentDirectory, showHidden)

		return directoryMsg(files)
	}
}

func readFileContent(file string) tea.Cmd {
	return func() tea.Msg {
		content := utils.ReadFileContent(file)

		return fileContentMsg(content)
	}
}

func createDir(dir string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		utils.CreateDirectory(dir)
		files := utils.GetDirectoryListing(constants.CurrentDirectory, showHidden)

		return directoryMsg(files)
	}
}

func createFile(name string, showHidden bool) tea.Cmd {
	return func() tea.Msg {
		utils.CreateFile(name)
		files := utils.GetDirectoryListing(constants.CurrentDirectory, showHidden)

		return directoryMsg(files)
	}
}
