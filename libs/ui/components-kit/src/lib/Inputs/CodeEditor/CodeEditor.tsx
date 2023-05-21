import { Label } from '../Label/Label';
import CodeMirror, { ReactCodeMirrorProps } from '@uiw/react-codemirror';
import { StreamLanguage } from '@codemirror/language';
import { shell } from '@codemirror/legacy-modes/mode/shell';

export type CodeEditorProps = Omit<
  ReactCodeMirrorProps,
  'theme' | 'setOptions'
> & {
  lang: 'sh';
  label: string;
};

export const CodeEditor = ({
  lang,
  label,
  ...editorProps
}: CodeEditorProps) => {
  return (
    <div className={'my-2'}>
      <Label>
        <span>{label}</span>
        <CodeMirror
          height="200px"
          extensions={[StreamLanguage.define(shell)]}
          {...editorProps}
        />
        {/*<AceEditor*/}
        {/*  className={'rounded-lg w-full h-64'}*/}
        {/*  mode={lang}*/}
        {/*  theme="github"*/}
        {/*  editorProps={{ $blockScrolling: true }}*/}
        {/*  showPrintMargin={true}*/}
        {/*  showGutter={true}*/}
        {/*  highlightActiveLine={true}*/}
        {/*  height={'200px'}*/}
        {/*  width={'100%'}*/}
        {/*  setOptions={{*/}
        {/*    enableBasicAutocompletion: true,*/}
        {/*    enableLiveAutocompletion: true,*/}
        {/*    enableSnippets: true,*/}
        {/*    showLineNumbers: true,*/}
        {/*    tabSize: 4,*/}
        {/*  }}*/}
        {/*  {...editorProps}*/}
        {/*/>*/}
      </Label>
    </div>
  );
};
